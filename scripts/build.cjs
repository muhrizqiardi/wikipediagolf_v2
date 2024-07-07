const esbuild = require("esbuild");
const pluginPostcss = require("esbuild-postcss");

async function build() {
  await esbuild.build({
    entryPoints: [
      "./client/src/index.js",
      "./client/src/surreal.js",
      "./client/src/styles.css",
    ],
    outdir: "./client/dist",
    bundle: true,
    minify: true,
    allowOverwrite: true,
    plugins: [pluginPostcss()],
  });
}
build().catch((error) => console.error(error));
