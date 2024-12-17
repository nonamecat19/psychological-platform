import { defineStore } from 'pinia'
import { ref } from 'vue'
import { api } from '@core/lib/axios.ts'

export const useTherapyGroupsStore = defineStore(
  'therapy-groups-default',
  () => {
    const data = ref()

    function setData(value: unknown) {
      data.value = value
    }

    async function init() {
      const response = await api.get('/therapy-group/my')
      setData(response.data)
    }

    return {
      data,
      setData,
      init,
    }
  },
)
