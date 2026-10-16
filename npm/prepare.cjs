const fs = require('node:fs');
const path = require('node:path');

const validateOnly = process.argv[2] === '--validate';
const tag = process.argv[validateOnly ? 3 : 2] || '';
const version = tag.replace(/^v/, '');
const number = '(?:0|[1-9][0-9]*)';
const identifier = '(?:0|[1-9][0-9]*|[0-9]*[A-Za-z-][0-9A-Za-z-]*)';
const buildSuffix = `(?:\\+[0-9A-Za-z-]+(?:\\.[0-9A-Za-z-]+)*)?`;
if (!new RegExp(`^${number}\\.${number}\\.${number}(?:-${identifier}(?:\\.${identifier})*)?${buildSuffix}$`).test(version)) {
  console.error('Specify VERSION as a release version, for example v0.1.0-beta.1.');
  process.exit(1);
}
if (validateOnly) process.exit(0);

const root = path.resolve(__dirname, '..');
const build = path.join(root, 'build', tag);
const output = path.join(build, 'npm');
const template = require('./package.json');
const optionalDependencies = {};
for (const [goos, os] of [['linux', 'linux'], ['darwin', 'darwin'], ['windows', 'win32']]) {
  for (const [goarch, cpu] of [['amd64', 'x64'], ['arm64', 'arm64']]) {
    const name = `@infituit/rss-cli-${os}-${cpu}`;
    const dir = path.join(output, `${os}-${cpu}`);
    const binary = `rss-cli${os === 'win32' ? '.exe' : ''}`;
    fs.mkdirSync(path.join(dir, 'bin'), { recursive: true });
    fs.copyFileSync(path.join(build, `rss-cli_${goos}_${goarch}${os === 'win32' ? '.exe' : ''}`), path.join(dir, 'bin', binary));
    fs.chmodSync(path.join(dir, 'bin', binary), 0o755);
    fs.writeFileSync(path.join(dir, 'package.json'), JSON.stringify({
      name, version, description: `${template.description} (${os}/${cpu})`,
      os: [os], cpu: [cpu], files: ['bin/'],
      repository: template.repository, publishConfig: template.publishConfig,
    }, null, 2) + '\n');
    optionalDependencies[name] = version;
  }
}

const main = path.join(output, 'rss-cli');
fs.mkdirSync(main, { recursive: true });
const { private: privatePackage, ...manifest } = template;
fs.writeFileSync(path.join(main, 'package.json'), JSON.stringify({
  ...manifest, version, optionalDependencies,
}, null, 2) + '\n');
fs.copyFileSync(path.join(__dirname, 'run.cjs'), path.join(main, 'run.cjs'));
fs.chmodSync(path.join(main, 'run.cjs'), 0o755);
console.log(`Prepared npm packages in ${output}`);
