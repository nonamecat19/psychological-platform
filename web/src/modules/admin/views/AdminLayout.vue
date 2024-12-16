<script setup lang="ts">
const links = [
  {
    title: 'Users',
    href: '/admin/users',
  },
]
</script>
<template>
  <main class="flex gap-2 h-full">
    <div class="w-80 bg-zinc-900">
      <div v-for="{ title, href } in links">
        <RouterLink :to="href">
          {{ title }}
        </RouterLink>
      </div>
    </div>
    <RouterView />
  </main>
</template>
