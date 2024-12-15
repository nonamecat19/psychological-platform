import { defineStore } from 'pinia'
import { ref } from 'vue'
import { jwtDecode } from 'jwt-decode'

export const useCurrentUserStore = defineStore('core-currentUser', () => {
  const data = ref()
  const isLoggedIn = ref<boolean>(false)

  function setData(value: unknown) {
    data.value = value
  }

  async function initFromToken() {
    const token = localStorage.getItem('token')
    if (token) {
      isLoggedIn.value = true
    } else {
      isLoggedIn.value = false
      return
    }
    try {
      return setData(jwtDecode(token))
    } catch {}
  }

  return {
    data,
    isLoggedIn,
    initFromToken,
  }
})
