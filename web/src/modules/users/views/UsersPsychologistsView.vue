<script setup lang="ts">
import { usePsychologistsStore } from '../store/usePsychologistsStore'
import { onMounted } from 'vue'
import { Button, Column, DataTable } from 'primevue'
import { useRouter } from 'vue-router'
import { api } from '@core/lib/axios'

const store = usePsychologistsStore()
const router = useRouter()

onMounted(() => {
  store.init()
})

async function contactPsychologist(id) {
  const { data } = await api.get<{ token: string }>(
    `/private-chat/get-by-user/${id}`,
  )
  await router.push(`/chat/${data}`)
}
</script>

<template>
  <DataTable :value="store.data">
    <Column field="Name" header="Name" />
    <Column field="Surname" header="Surname" />
    <Column field="Email" header="Email" />
    <Column field="Name" header="Name" />
    <Column>
      <template #body="slotProps">
        <Button @click="contactPsychologist(slotProps.data.ID)">Contact</Button>
      </template>
    </Column>
  </DataTable>
</template>
