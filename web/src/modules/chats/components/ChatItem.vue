<script setup lang="ts">
import { defineProps, defineEmits } from 'vue'

const props = defineProps({
  name: {
    type: String,
    default: 'Anonymous',
  },
})

const emits = defineEmits(['click'])

function handleClick() {
  emits('click')
}
</script>

<template>
  <button
    class="flex items-center gap-2 p-2 hover:bg-gray-900 duration-200 w-full"
    @click="handleClick"
  >
    <div
      class="size-10 rounded-full bg-gray-700 flex justify-center items-center"
    >
      {{ props.name?.substring(0, 1) }}
    </div>
    <div>
      {{ props.name }}
    </div>
  </button>
</template>

<style scoped></style>
