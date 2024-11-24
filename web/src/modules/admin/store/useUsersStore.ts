import { defineStore } from 'pinia'
import { ref } from 'vue'
import { api } from '@core/lib/axios.ts'

export const useUsersStore = defineStore('admin-users', () => {
  const data = ref()

  function setData(value: unknown) {
    data.value = value
  }

  async function init() {
    const response = await api.get('/users')
    setData(response.data)
  }

  return {
    data,
    setData,
    init,
  }
})
