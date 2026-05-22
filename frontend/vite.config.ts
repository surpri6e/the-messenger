import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import basicSsl from "@vitejs/plugin-basic-ssl";

import path from "path";

// https://vite.dev/config/
export default defineConfig({
  plugins: [react(), basicSsl()], // basicSsl()
  resolve: {
    alias: {
      "@": path.resolve(__dirname, "./src"),
      "@styles": path.resolve(__dirname, "./src/utils/styles"),
      "@images": path.resolve(__dirname, "./src/utils/images"),
      "@components": path.resolve(__dirname, "./src/components"),
      "@appTypes": path.resolve(__dirname, "./src/utils/types"),
      "@stores": path.resolve(__dirname, "./src/stores"),
      "@pages": path.resolve(__dirname, "./src/pages"),
      "@functionals": path.resolve(__dirname, "./src/utils/functionals"),
      "@hooks": path.resolve(__dirname, "./src/hooks"),
      "@api": path.resolve(__dirname, "./src/api"),
      "@constants": path.resolve(__dirname, "./src/utils/constants"),
    },
  },
});
