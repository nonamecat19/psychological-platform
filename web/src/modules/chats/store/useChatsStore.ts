import { defineStore } from 'pinia'
import { ref } from 'vue'
import { api } from '@core/lib/axios.ts'

export const useChatsStore = defineStore('chats-default', () => {
  const data = ref()

  function setData(value: unknown) {
    data.value = value
  }

  async function init() {
    const response = await api.get('/private-chat/my')
    setData(response.data)
  }

  return {
    data,
    setData,
    init,
  }
})
