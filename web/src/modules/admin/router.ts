export const adminRouter = {
  path: '/admin',
  component: () => import('./views/AdminLayout.vue'),
  children: [
    {
      path: 'users',
      component: () => import('./views/AdminUsersView.vue'),
    },
  ],
}
