<script setup lang="ts">
import { onMounted, ref } from 'vue'
import { useTherapyGroupsStore } from '../store/therapyGroups.ts'
import ChatItem from '../components/ChatItem.vue'
import { useRouter } from 'vue-router'
import Button from 'primevue/button'
import { useCurrentTherapyGroupStore } from '../store/currentTherapyGroup.ts'
import Dialog from 'primevue/dialog'
import { useCurrentUserStore } from '@core/stores/useCurrentUserStore.ts'
import AddTherapyGroupForm from '../components/AddTherapyGroupForm.vue'

const groupsStore = useTherapyGroupsStore()
const currentGroupStore = useCurrentTherapyGroupStore()
const currentUserStore = useCurrentUserStore()

onMounted(() => {
  groupsStore.init()
})

const router = useRouter()

function selectChat(id: number) {
  currentGroupStore.setCurrentTherapyGroupId(id)
  currentGroupStore.initCurrentTherapyGroupId()
  router.push(`/therapy-groups/${id}`)
}

const visible = ref<boolean>(false)
</script>
<template>
  <main class="flex h-full w-full">
    <Dialog
      v-model:visible="visible"
      modal
      header="Add new therapy group"
      :style="{ width: '25rem' }"
    >
      <AddTherapyGroupForm
        @hide="visible = false"
        @update-data="groupsStore.init()"
      />
    </Dialog>

    <div class="border-r w-64 h-full overflow-y-auto">
      <div class="p-2">
        <Button
          @click="visible = true"
          v-if="currentUserStore.data.role === 'psychologist'"
          class="w-full"
        >
          Add new
        </Button>
      </div>

      <ChatItem
        v-for="(chat, index) in groupsStore.data"
        :key="index"
        :name="chat.Name"
        :selected="currentGroupStore.currentTherapyGroupId === chat.ID"
        @click="selectChat(chat.ID)"
      />
    </div>
    <div class="w-full">
      <RouterView />
    </div>
  </main>
</template>
