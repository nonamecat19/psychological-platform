import SupportView from "./views/SupportView.vue";

export const supportRouter = {
  path: '/support',
  children: [
    {
      path: '',
      component: SupportView
    }
  ]
}
