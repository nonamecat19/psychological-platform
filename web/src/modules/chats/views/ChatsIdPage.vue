<script setup lang="ts">
import { useRoute } from 'vue-router'
import { computed, onMounted, ref } from 'vue'
import { useCurrentChatStore } from '@chats/store/currentChat'
import InputText from 'primevue/inputtext'
import Button from 'primevue/button'
import { useCurrentChatSubscription } from '@chats/composables/useCurrentChatSubscription'
import ChatMessage from '@chats/components/ChatMessage.vue'
import { useCurrentUserStore } from '@core/stores/useCurrentUserStore.ts'

const route = useRoute()

const currentChatStore = useCurrentChatStore()

const id = computed(() => route.params.id)

onMounted(() => {
  currentChatStore.setCurrentChatId(Number(id.value))
  currentChatStore.initCurrentChatId()
})

const { sendMessage } = useCurrentChatSubscription(+id.value)

async function handleSendMessage() {
  sendMessage(inputValue.value)
  inputValue.value = ''
}

const currentUserStore = useCurrentUserStore()

const inputValue = ref<string>('')
</script>
<template>
  <div class="h-full w-full flex flex-col">
    <div class="grow p-4 flex flex-col gap-2">
      <ChatMessage
        v-for="message of currentChatStore?.data?.Messages"
        :text="message.Content"
        :is-mine="message.UserID === currentUserStore.data.ID"
      />
    </div>
    <div class="p-2 pt-0 w-full flex gap-2">
      <InputText v-model="inputValue" class="w-full" />
      <Button @click="handleSendMessage">
        <i class="pi pi-send" />
      </Button>
    </div>
  </div>
</template>
