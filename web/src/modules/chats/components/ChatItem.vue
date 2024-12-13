<script setup lang="ts">
defineProps({
  name: {
    type: String,
    default: 'Anonymous',
  },
  selected: {
    type: Boolean,
    default: false,
  },
})

const emits = defineEmits(['click'])

function handleClick() {
  emits('click')
}
</script>

<template>
  <button
    :class="[
      'flex items-center gap-2 p-2 hover:bg-gray-900 duration-200 w-full',
      selected && 'bg-white hover:bg-gray-50 text-gray-900',
    ]"
    @click="handleClick"
  >
    <div
      :class="[
        'size-10 rounded-full bg-gray-700 flex justify-center items-center',
        selected && 'text-white',
      ]"
    >
      {{ name?.substring(0, 1) }}
    </div>
    <div>
      {{ name }}
    </div>
  </button>
</template>
