const esbuild = require("esbuild");
const pluginPostcss = require("esbuild-postcss");

async function build() {
  await esbuild.build({
    entryPoints: ["./client/src/index.ts", "./client/src/styles.css"],
    sourcemap: true,
    outdir: "./client/dist",
    bundle: true,
    allowOverwrite: true,
    plugins: [pluginPostcss()],
  });
}
build().catch((error) => console.error(error));
