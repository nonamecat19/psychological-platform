<script setup lang="ts">
import { useProfileStore } from '../store/profile.ts'
import { onMounted } from 'vue'

const profileStore = useProfileStore()

onMounted(() => {
  profileStore.init()
})
</script>

<template>
  <h2>My info</h2>
  <h3>Name: {{ profileStore.data?.Name }}</h3>
  <h3>Surname: {{ profileStore.data?.Surname }}</h3>
  <h3>Email: {{ profileStore.data?.Email }}</h3>
  <h3>Anonymous?: {{ profileStore.data?.IsAnonymous ? 'Yes' : 'No' }}</h3>
  <h3>Role: {{ profileStore.data?.Role }}</h3>
  <h3>Created at: {{ profileStore.data?.CreatedAt }}</h3>
  <h3>Updated at: {{ profileStore.data?.UpdatedAt }}</h3>
</template>
