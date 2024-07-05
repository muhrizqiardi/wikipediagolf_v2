const esbuild = require("esbuild");
const pluginPostcss = require("esbuild-plugin-postcss2");

async function build() {
  await esbuild.build({
    entryPoints: ["./client/src/index.ts", "./client/src/styles.css"],
    outdir: "./client/dist",
    bundle: true,
    minify: true,
    allowOverwrite: true,
    plugins: [
      pluginPostcss.default({
        plugins: [require("autoprefixer"), require("tailwindcss")],
      }),
    ],
  });
}
build().catch((error) => console.error(error));
