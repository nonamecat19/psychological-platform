<script setup>
import { usePsychologistsStore } from '../store/usePsychologistsStore'
import { onMounted } from 'vue'
import { Button, Column, DataTable } from 'primevue'

const store = usePsychologistsStore()

onMounted(() => {
  store.init()
})


function contactPsychologist(id) {
  console.log('Contact psychologist with ID:', id)
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

<style scoped></style>
