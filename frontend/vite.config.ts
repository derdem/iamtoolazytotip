import { defineConfig } from 'vite';
import solidPlugin from 'vite-plugin-solid';
import devtools from "solid-devtools/vite"

export default defineConfig({
  plugins: [
    devtools({
      /* features options - all disabled by default */
      autoname: true, // e.g. enable autoname
      locator: {
        targetIDE: 'vscode',
        componentLocation: true,
        jsxLocation: true,
      },
    }),
    solidPlugin()
  ],
  server: {
    port: 3000,
  },
  build: {
    target: 'esnext',
  },
  preview: {
    port: 4174,
  },
});
