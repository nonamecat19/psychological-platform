import { createRouter, createWebHistory } from 'vue-router'
import { chatRouter } from '@chats'
import { supportRouter } from '@support'
import { adminRouter } from '../modules/admin'
import { authRouter } from '../modules/auth'
import { usersRouter } from '../modules/users'
import { profileRouter } from '../modules/profile'
import { therapyGroupRouter } from '../modules/therapyGroups'
import MainLayout from '@app/layouts/MainLayout.vue'
import { useCurrentUserStore } from '@core/stores/useCurrentUserStore.ts'

const routes = [
  chatRouter,
  supportRouter,
  adminRouter,
  authRouter,
  usersRouter,
  profileRouter,
  therapyGroupRouter,
]

export const appRouter = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/',
      component: MainLayout,
      children: routes,
    },
  ],
})

appRouter.beforeEach((_to, _from, next) => {
  const currentUserStore = useCurrentUserStore()
  void currentUserStore.initFromToken()
  next()
})
