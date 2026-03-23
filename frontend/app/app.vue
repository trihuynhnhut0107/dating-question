<script setup lang="ts">
useHead({
  meta: [{ name: 'viewport', content: 'width=device-width, initial-scale=1' }],
  link: [{ rel: 'icon', href: '/favicon.ico' }],
  htmlAttrs: { lang: 'en' }
})

useSeoMeta({
  title: 'DateDeck – Shuffle your conversations',
  description: 'Draw random dating questions to spark real conversations.'
})

const links = [
  { label: 'Home', to: '/', icon: 'i-lucide-home' },
  { label: 'Questions', to: '/questions', icon: 'i-lucide-help-circle' },
  { label: 'Categories', to: '/categories', icon: 'i-lucide-layers' },
  { label: 'Sessions', to: '/sessions', icon: 'i-lucide-heart' }
]
</script>

<template>
  <UApp>
    <UHeader>
      <template #left>
        <NuxtLink to="/" class="font-bold text-xl text-primary flex items-center gap-2">
          <UIcon name="i-lucide-heart" class="text-primary" />
          DateDeck
        </NuxtLink>
      </template>
      <UNavigationMenu :items="links" />
      <template #right>
        <UColorModeButton />
      </template>
    </UHeader>

    <UMain>
      <NuxtPage />
    </UMain>

    <UFooter>
      <template #left>
        <p class="text-muted text-sm">© {{ new Date().getFullYear() }} DateDeck</p>
      </template>
    </UFooter>
  </UApp>
</template>
