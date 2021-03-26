<template>
  <div>

    <div v-if="post">
         <div class="album py-5 bg-light" style="margin-top:50px;">
              <div class="container">
                  <h1>{{ post.Title }}</h1>
                  <p>{{ post.Body }}</p>
                  <div class="d-flex justify-content-between align-items-center">

                      <smal class="text-muted">Written by {{ post.UserId }}</smal>
                  </div>
              </div>

          </div>
          <div class="container"   style="margin-top:50px;">
              <div class="text-muted py-2">
                  <button class="btn btn-sm btn-outline-secondary" style="margin: auto">Add coment</button>
              </div>
          </div>
       </div>
  </div>
</template>

<script>
  import axios from 'axios'
export default {
  props: [
    'id',
     ],
  data: function(){
    return {
      post: null
    }
  },
  watch: {
    '$route' (to,from){
      this.getData()
    }
  },
  methods: {
    getData (){
     // this.id
     var that=this
     // axios.get ('https://jsonplaceholder.typicode.com/posts/'+this.id)
      axios.get ('http://localhost:8000/posts/'+this.id)
        .then(function (response) {
        that.post = response.data
        })

    }

  },
  mounted: function (){
    this.getData()

  }

}
</script>

