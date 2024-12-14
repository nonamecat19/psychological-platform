<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useMe } from '@core/composables/useMe.ts'

const isLogged = !!localStorage.getItem('token')

const { getDecodedToken } = useMe()

const me = getDecodedToken()

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
        <RouterLink v-if="me?.role === 'user'" to="/users/psychologists"
          >Psychologists
        </RouterLink>
      </div>
      {{ me }}
      <button v-if="isLogged" @click="handleLogout">Logout</button>
      <RouterLink v-else to="/auth/login">Login</RouterLink>
    </header>
    <hr />
    <div class="h-[calc(100vh-57px)]">
      <RouterView />
    </div>
  </main>
</template>
