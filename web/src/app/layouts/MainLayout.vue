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
      <div class="flex gap-2">
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
      <button v-if="currentUserStore.isLoggedIn" @click="handleLogout">
        Logout
      </button>
      <RouterLink v-else to="/auth/login">Login</RouterLink>
    </header>
    <hr />
    <div class="h-[calc(100vh-57px)]">
      <RouterView />
    </div>
  </main>
</template>
