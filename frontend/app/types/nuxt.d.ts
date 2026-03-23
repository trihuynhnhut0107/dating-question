// Augment Nuxt's runtime config so config.public.apiBase is typed as string
// without needing an `as string` cast anywhere.
declare module 'nuxt/schema' {
  interface PublicRuntimeConfig {
    apiBase: string
  }
}

export {}
