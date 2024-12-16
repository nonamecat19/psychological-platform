<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useCurrentUserStore } from '@core/stores/useCurrentUserStore.ts'

const currentUserStore = useCurrentUserStore()
const router = useRouter()

function handleLogout() {
  localStorage.removeItem('token')
  router.push('/auth/login')
}
</script>

<template>
  <main>
    <header class="w-full flex justify-between p-4 overflow-hidden">
      <div class="flex gap-8">
        <RouterLink to="/">Home</RouterLink>
        <RouterLink to="/chat">Chats</RouterLink>
        <RouterLink
          v-if="currentUserStore.data?.role === 'user'"
          to="/users/psychologists"
          >Psychologists
        </RouterLink>
        <RouterLink v-if="currentUserStore.data" to="/profile"
          >Profile</RouterLink
        >
      </div>
      <div class="flex gap-8">
        <RouterLink to="/admin">Admin</RouterLink>
        <button v-if="currentUserStore.isLoggedIn" @click="handleLogout">
          Logout
        </button>
        <RouterLink v-else to="/auth/login">Login</RouterLink>
      </div>
    </header>
    <hr />
    <div class="h-[calc(100vh-57px)]">
      <RouterView />
    </div>
  </main>
</template>
