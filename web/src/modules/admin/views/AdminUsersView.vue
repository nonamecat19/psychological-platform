<script setup lang="ts">
import { onMounted, ref } from 'vue'
import DataTable from 'primevue/datatable'
import Column from 'primevue/column'
import { useAdminUsersStore } from '../store/useAdminUsersStore.ts'
import Button from 'primevue/button'
import Dialog from 'primevue/dialog'
import AddUserForm from '../components/AddUserForm.vue'

const store = useAdminUsersStore()

onMounted(() => {
  store.init()
})

const visible = ref<boolean>(false)
</script>

<template>
  <main class="w-full flex flex-col gap-2">
    <Button label="Add" @click="visible = true" class="w-20" />

    <Dialog
      v-model:visible="visible"
      modal
      header="Edit Profile"
      :style="{ width: '25rem' }"
    >
      <AddUserForm @hide="visible = false" @update-data="store.init()" />
    </Dialog>
    <DataTable
      v-if="store.data"
      :value="store.data"
      class="w-full"
      size="large"
      paginator
      :rows="5"
      :rowsPerPageOptions="[5, 10, 20, 50]"
    >
      <Column field="ID" header="ID" />
      <Column field="Name" header="Name" />
      <Column field="Surname" header="Surname" />
      <Column field="Email" header="Email" />
      <Column field="IsAnonymous" header="Is anonymous?" />
      <Column field="Role" header="Role" />
    </DataTable>
  </main>
</template>
