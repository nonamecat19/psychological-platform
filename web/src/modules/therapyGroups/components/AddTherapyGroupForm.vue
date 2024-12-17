<script setup lang="ts">
import Button from 'primevue/button'
import InputText from 'primevue/inputtext'
import * as yup from 'yup'
import { useForm } from 'vee-validate'
import { toTypedSchema } from '@vee-validate/yup'
import { useToast } from 'primevue/usetoast'
import { api } from '@core/lib/axios.ts'
import { AxiosError } from 'axios'

const emits = defineEmits(['hide', 'updateData'])

function handleHide() {
  emits('hide')
}

const authSchema = yup.object({
  name: yup.string().required(),
  description: yup.string().required(),
})

const { defineField, handleSubmit, errors } = useForm({
  validationSchema: toTypedSchema(authSchema),
})

const [name] = defineField('name')
const [description] = defineField('description')

const toast = useToast()

const onSubmit = handleSubmit(async ({ name, description }) => {
  try {
    await api.post('/therapy-group', {
      name,
      description,
    })
    emits('updateData')
    emits('hide')
  } catch (error) {
    if (error instanceof AxiosError) {
      toast.add({
        severity: 'error',
        summary: error.response?.data.message,
        life: 2000,
      })
    }
  }
})
</script>

<template>
  <form @submit="onSubmit">
    <div>
      <label class="block mt-5 mb-1" for="name">Name</label>
      <InputText
        v-model="name"
        aria-describedby="name-help"
        type="text"
        :class="{ 'p-invalid': errors.name }"
      />
      <small id="name-help" class="block">{{ errors.name }}</small>
    </div>

    <div>
      <label class="block mt-5 mb-1" for="description">Description</label>
      <InputText
        v-model="description"
        aria-describedby="description-help"
        type="text"
        :class="{ 'p-invalid': errors.description }"
      />
      <small id="description-help" class="block">{{
        errors.description
      }}</small>
    </div>

    <div class="flex gap-2">
      <Button type="button" class="mt-5" @click="handleHide">Cancel</Button>
      <Button type="submit" class="mt-5">Submit</Button>
    </div>
  </form>
</template>
