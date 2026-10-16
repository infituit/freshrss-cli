const assert = require('node:assert/strict');
const { execFileSync, spawnSync } = require('node:child_process');
const fs = require('node:fs');
const os = require('node:os');
const path = require('node:path');
const { test } = require('node:test');

test('packed CLI installs without scripts and preserves native output', () => {
  const tag = process.env.VERSION;
  assert.ok(tag, 'Set VERSION to the version passed to make npm-packages');
  const packages = path.resolve(__dirname, '..', 'build', tag, 'npm');
  const temp = fs.mkdtempSync(path.join(os.tmpdir(), 'rss-cli-npm-'));
  const npm = process.platform === 'win32' ? 'npm.cmd' : 'npm';
  const platform = `${process.platform}-${process.arch}`;
  try {
    const tarballs = [];
    for (const name of ['rss-cli', platform]) {
      const packed = JSON.parse(execFileSync(npm, [
        'pack', path.join(packages, name), '--json', '--ignore-scripts',
        '--pack-destination', temp,
      ], { encoding: 'utf8', cwd: temp }));
      const result = Array.isArray(packed) ? packed[0] : Object.values(packed)[0];
      tarballs.push(path.join(temp, result.filename));
    }
    execFileSync(npm, [
      'install', '--prefix', temp, '--offline', '--ignore-scripts',
      '--no-audit', '--no-fund', ...tarballs,
    ], { cwd: temp, stdio: 'pipe' });

    const main = path.join(temp, 'node_modules', '@infituit', 'rss-cli');
    const launcher = path.join(main, 'run.cjs');
    const native = path.join(temp, 'node_modules', '@infituit',
      `rss-cli-${platform}`, 'bin', `rss-cli${process.platform === 'win32' ? '.exe' : ''}`);
    const manifest = JSON.parse(fs.readFileSync(path.join(main, 'package.json')));
    assert.equal(Object.keys(manifest.optionalDependencies).length, 6);
    assert.ok(Object.values(manifest.optionalDependencies).every((v) => v === tag.replace(/^v/, '')));
    for (const args of [['version'], ['invalid-command-with-spaces']]) {
      const direct = spawnSync(native, args, { encoding: 'utf8' });
      const wrapped = spawnSync(process.execPath, [launcher, ...args], { encoding: 'utf8' });
      assert.equal(wrapped.status, direct.status);
      assert.equal(wrapped.stdout, direct.stdout);
      assert.equal(wrapped.stderr, direct.stderr);
    }
    const version = spawnSync(process.execPath, [launcher, 'version'], { encoding: 'utf8' });
    assert.equal(JSON.parse(version.stdout).version, tag);

    fs.rmSync(path.dirname(path.dirname(native)), { recursive: true });
    const missing = spawnSync(process.execPath, [launcher, 'version'], { encoding: 'utf8' });
    assert.equal(missing.status, 1);
    assert.equal(missing.stdout, '');
    assert.match(JSON.parse(missing.stderr).error.message, /--include=optional/);
  } finally {
    fs.rmSync(temp, { recursive: true, force: true });
  }
});
