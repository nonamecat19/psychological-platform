<script setup lang="ts">
import * as yup from 'yup'
import { useForm } from 'vee-validate'
import { toTypedSchema } from '@vee-validate/yup'
import { useToast } from 'primevue/usetoast'
import { AxiosError } from 'axios'
import { api } from '@core/lib/axios.ts'
import { useRouter } from 'vue-router'
import { Button, Select, InputText } from 'primevue'

const roles = [
  { name: 'Client', code: 'client' },
  { name: 'Psychologist', code: 'psychologist' },
]

const isAnonymousOptions = [
  { name: 'Yes', code: 'yes' },
  { name: 'No', code: 'no' },
]

const authSchema = yup.object({
  email: yup.string().required(),
  password: yup.string().required(),
  name: yup.string().required(),
  surname: yup.string().required(),
  isAnonymous: yup.object({ name: yup.string(), code: yup.string() }),
  role: yup.object({ name: yup.string(), code: yup.string() }).required(),
})

const { defineField, handleSubmit, errors } = useForm({
  validationSchema: toTypedSchema(authSchema),
})

const [email] = defineField('email')
const [password] = defineField('password')
const [name] = defineField('name')
const [surname] = defineField('surname')
const [isAnonymous] = defineField('isAnonymous')
const [role] = defineField('role')

const toast = useToast()
const router = useRouter()

const onSubmit = handleSubmit(
  async ({ email, password, name, surname, isAnonymous, role }) => {
    try {
      const { data } = await api.post<{ token: string }>('/auth/register', {
        email,
        password,
        name,
        surname,
        isAnonymous: isAnonymous.code === 'yes',
        role: role.code,
      })

      localStorage.setItem('token', data.token)
      await router.push('/')
    } catch (error) {
      if (error instanceof AxiosError) {
        toast.add({
          severity: 'error',
          summary: error.response?.data.message,
          life: 2000,
        })
      }
    }
  },
)
</script>

<template>
  <div>
    <h1 class="text-center">Register</h1>
    <div>
      <form @submit="onSubmit">
        <div>
          <label class="block mt-5 mb-1" for="email">Email</label>
          <InputText
            v-model="email"
            aria-describedby="email-help"
            type="email"
            :class="{ 'p-invalid': errors.email }"
          />
          <small id="email-help" class="block">{{ errors.email }}</small>
        </div>

        <div>
          <label class="block mt-5 mb-1" for="password">Password</label>
          <InputText
            v-model="password"
            aria-describedby="password-help"
            type="password"
            :class="{ 'p-invalid': errors.password }"
          />
          <small id="password-help" class="block">{{ errors.password }}</small>
        </div>

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
          <label class="block mt-5 mb-1" for="surname">Surname</label>
          <InputText
            v-model="surname"
            aria-describedby="surname-help"
            type="text"
            :class="{ 'p-invalid': errors.surname }"
          />
          <small id="surname-help" class="block">{{ errors.surname }}</small>
        </div>

        <div>
          <label class="block mt-5 mb-1" for="isAnonymous">Anonymous</label>
          <Select
            v-model="isAnonymous"
            :options="isAnonymousOptions"
            option-label="name"
            aria-describedby="isAnonymous-help"
            :class="{ 'p-invalid': errors.isAnonymous }"
          />
          <small id="isAnonymous-help" class="block">{{
            errors.isAnonymous
          }}</small>
        </div>

        <div>
          <label class="block mt-5 mb-1" for="role">Role</label>
          <Select
            v-model="role"
            :options="roles"
            option-label="name"
            aria-describedby="role-help"
            :class="{ 'p-invalid': errors.role }"
          />
          <small id="role-help" class="block">{{ errors.role }}</small>
        </div>

        <Button type="submit" class="mt-5">Submit</Button>
      </form>
    </div>
  </div>
</template>
