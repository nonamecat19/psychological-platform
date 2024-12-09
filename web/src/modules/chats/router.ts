export const chatRouter = {
  path: '/chat',
  component: () => import('./views/ChatsLayout.vue'),
  children: [
    {
      path: ':id',
      component: () => import('./views/ChatsIdPage.vue'),
    },
  ],
}
