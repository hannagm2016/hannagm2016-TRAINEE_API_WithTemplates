import Vue from 'vue'
import Router from 'vue-router'
import Index from '@/components/Index'
import Edit from '@/components/Edit'
import Create from '@/components/Edit'
import Post from '@/components/Post'
import Authorization from '@/components/Authorization'
import Registration from '@/components/Registration'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      name: 'Index',
      component: Index
    },
    {
      path: '/post/:id',
      name: 'Post',
      props: true,
      component: Post
    },
    {
      path: '/create',
      name: 'create',
      component: Create
    },
    {
      path: '/edit/:id',
      name: 'Edit',
        props: true,
      component: Edit
    },
      {
          path: '/authorization',
          name: 'Authorization',
            props: true,
          component: Authorization
        },
         {
                  path: '/registration',
                  name: 'Registration',
                    props: true,
                  component: Registration
                }

  ]
})
