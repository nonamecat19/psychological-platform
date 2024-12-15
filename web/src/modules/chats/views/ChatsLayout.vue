<script setup lang="ts">
import { onMounted } from 'vue'
import { useChatsStore } from '@chats/store/chats.ts'
import ChatItem from '@chats/components/ChatItem.vue'
import { useRouter } from 'vue-router'
import { useCurrentChatStore } from '@chats/store/currentChat.ts'

const chatsStore = useChatsStore()
const currentChatStore = useCurrentChatStore()

onMounted(() => {
  chatsStore.init()
})

const router = useRouter()

function selectChat(id: number) {
  currentChatStore.setCurrentChatId(id)
  currentChatStore.initCurrentChatId()
  router.push(`/chat/${id}`)
}
</script>
<template>
  <main class="flex h-full w-full">
    <div class="border-r w-64 h-full overflow-y-auto">
      <ChatItem
        v-for="(chat, index) in chatsStore.data"
        :key="index"
        :name="chat.Specialist.Name ?? 'Anonymous'"
        :selected="currentChatStore.currentChatId === chat.ID"
        @click="selectChat(chat.ID)"
      />
    </div>
    <div class="w-full">
      <RouterView />
    </div>
  </main>
</template>
