import { defineStore } from 'pinia'
import { ref } from 'vue'
import { api } from '@core/lib/axios.ts'

export const usePsychologistsStore = defineStore('users-psychologists', () => {
  const data = ref()

  function setData(value: unknown) {
    data.value = value
  }

  async function init() {
    const response = await api.get('/users/psychologists')
    setData(response.data)
  }

  return {
    data,
    setData,
    init,
  }
})
