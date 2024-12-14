import { defineStore } from 'pinia'
import { ref } from 'vue'
import { api } from '@core/lib/axios.ts'

export const useCurrentChatStore = defineStore('chats-current', () => {
  const data = ref()
  const currentChatId = ref<number | null>(null)

  function setData(value: unknown) {
    data.value = value
  }

  function setCurrentChatId(chatId: number) {
    currentChatId.value = chatId
  }

  async function initCurrentChatId() {
    const response = await api.get(`/private-chat/${currentChatId.value}`)
    setData(response.data)
  }

  function addMessage(message: unknown) {
    data.value.Messages.push(message)
  }

  return {
    data,
    currentChatId,
    setCurrentChatId,
    initCurrentChatId,
    addMessage,
  }
})
