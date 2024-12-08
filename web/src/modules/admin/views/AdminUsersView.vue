<script setup lang="ts">
import { useUsersStore } from '../store/useUsersStore.ts'
import { onMounted } from 'vue'

const store = useUsersStore()

onMounted(() => {
  store.init()
})
</script>

<template>
  <div>{{ JSON.stringify(store.data) }}</div>
</template>
