import { defineStore } from 'pinia'
import { ref } from 'vue'
import { api } from '@core/lib/axios.ts'

export const useCurrentTherapyGroupStore = defineStore(
  'therapy-groups-current',
  () => {
    const data = ref()
    const currentTherapyGroupId = ref<number | null>(null)

    function setData(value: unknown) {
      data.value = value
    }

    function setCurrentTherapyGroupId(chatId: number) {
      currentTherapyGroupId.value = chatId
    }

    async function initCurrentTherapyGroupId() {
      const response = await api.get(
        `/therapy-group/${currentTherapyGroupId.value}`,
      )
      setData(response.data)
    }

    function addMessage(message: unknown) {
      data.value.Messages.push(message)
    }

    return {
      data,
      currentTherapyGroupId,
      setCurrentTherapyGroupId,
      initCurrentTherapyGroupId,
      addMessage,
    }
  },
)
