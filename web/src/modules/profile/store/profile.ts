import { defineStore } from 'pinia'
import { ref } from 'vue'
import { api } from '@core/lib/axios.ts'

export const useProfileStore = defineStore('profile-default', () => {
  const data = ref()

  function setData(value: unknown) {
    data.value = value
  }

  async function init() {
    const response = await api.get('/users/me')
    console.log({ response })
    setData(response.data)
  }

  return {
    data,
    setData,
    init,
  }
})
