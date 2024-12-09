<script setup lang="ts">
import { onMounted } from 'vue'
import { useChatsStore } from '@chats/store/useChatsStore'
import ChatItem from '@chats/components/ChatItem.vue'
import { useRouter } from 'vue-router'

const store = useChatsStore()

onMounted(() => {
  store.init()
})

const router = useRouter()

function selectChat(id) {
  router.push(`/chat/${id}`)
}
</script>
<template>
  <main class="flex h-screen">
    <div class="border-r w-64 h-screen overflow-y-auto">
      <ChatItem
        v-for="(chat, index) in store.data"
        :key="index"
        :name="chat.Specialist.Name ?? 'Anonymous'"
        @click="selectChat(chat.ID)"
      />
    </div>
    <div>
      <RouterView />
    </div>
  </main>
</template>
