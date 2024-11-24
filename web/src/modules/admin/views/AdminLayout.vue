<script setup lang="ts">
const links = [
  {
    title: 'Users',
    href: '/admin/users',
  },
  {
    title: 'Therapy groups',
    href: '/admin/therapyGroups',
  },
]
</script>
<template>
  <main class="flex gap-5">
    <div class="w-80">
      <div v-for="{ title, href } in links">
        <RouterLink :to="href">
          {{ title }}
        </RouterLink>
      </div>
    </div>
    <RouterView />
  </main>
</template>
