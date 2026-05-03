import esbuild from "esbuild";

export default function (eleventyConfig) {
  eleventyConfig.on("eleventy.before", async () => {
    await esbuild.build({
      entryPoints: ["src/game-entry.ts"],
      bundle: true,
      outfile: "_site/game.js",
      minify: process.env.NODE_ENV === "production",
    });
  });

  eleventyConfig.addPassthroughCopy({ public: "." });

  return {
    dir: {
      input: "src",
      includes: "_includes",
      output: "_site",
    },
  };
}
