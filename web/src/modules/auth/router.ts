export const authRouter = {
  path: '/auth',
  component: () => import('./layouts/AuthLayout.vue'),
  children: [
    {
      path: 'login',
      component: () => import('./views/AuthLoginView.vue'),
    },
    {
      path: 'register',
      component: () => import('./views/AuthRegisterView.vue'),
    },
  ],
}
