#!/usr/bin/env node
const { spawn } = require('node:child_process');

function fail(message) {
  console.error(JSON.stringify({ error: { code: 'npm_launcher_error', message } }));
  process.exit(1);
}

const os = process.platform;
const cpu = process.arch;
if (!['linux', 'darwin', 'win32'].includes(os) || !['x64', 'arm64'].includes(cpu)) {
  fail(`Unsupported platform: ${os}/${cpu}.`);
}

const name = `@infituit/rss-cli-${os}-${cpu}`;
const version = require('./package.json').version;
let binary;
try {
  const manifest = require(`${name}/package.json`);
  if (manifest.version !== version) {
    fail(`Platform package ${name} must match version ${version}.`);
  }
  binary = require.resolve(`${name}/bin/rss-cli${os === 'win32' ? '.exe' : ''}`);
} catch {
  fail(`Missing platform package ${name}@${version}. Reinstall @infituit/rss-cli@${version} with optional dependencies enabled (--include=optional).`);
}

const child = spawn(binary, process.argv.slice(2), { stdio: 'inherit' });
const handlers = new Map();
for (const signal of ['SIGINT', 'SIGTERM']) {
  const handler = () => child.kill(signal);
  handlers.set(signal, handler);
  process.on(signal, handler);
}
child.on('error', (error) => fail(`Cannot start ${name}: ${error.message}`));
child.on('exit', (code, signal) => {
  for (const [name, handler] of handlers) process.removeListener(name, handler);
  if (signal) process.kill(process.pid, signal);
  else process.exit(code ?? 1);
});
