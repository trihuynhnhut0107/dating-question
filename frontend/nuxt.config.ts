// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  compatibilityDate: "2025-07-15",
  css: ["~/assets/css/main.css"],
  future: {
    compatibilityVersion: 4,
  },
  devtools: { enabled: true },
  modules: ["@vueuse/nuxt"],
  runtimeConfig: {
    apiSecret: "",
    public: {
      apiBaseUrl: "http://localhost:8080/api/v1",
    },
  },
});
