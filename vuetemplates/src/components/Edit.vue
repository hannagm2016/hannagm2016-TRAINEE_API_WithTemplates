<template>
<div>
  <h1>Edit {{id}}</h1>

    <div class="row py-lg-5">
                <div class="col-lg-6 col-md-8 mx-auto">
                    <form @submit.prevent="savePost">
                        <input type="hidden" name="id" :value="posts.Id" />
                        <input type="hidden" name="userid" :value="posts.UserId" />
                        <input v-model="posts.Title" type="text" id="title" placeholder="Title of your post" class="form-control"><br>
                        <textarea name="body"  v-model="posts.Body" placeholder="Body of your post" class="form-control"></textarea><br>
                        <button type="submit" class="btn btn-primary my-2" >Save</button>
                       </form>
                    <p>{{posts.Id}}</p>
                </div>

            </div>

  </div>
</template>

<script>
import axios from 'axios'
export default {
  props: [
    'id'
     ],
  data: () => ({
      posts: {
        Id: '',
        Title: '',
        Body: '',
        UserId: ''
      }
    }),
  watch: {
    '$route' (to,from){
      this.getData()
    }
  },
  methods: {
   savePost() {
   axios.post (`http://localhost:8000/savePost`, {
                  Id: this.posts.Id,
                  UserId: this.posts.UserId,
                  Title: this.posts.Title,
                  Body: this.posts.Body,
             })
              .then (response => {
             console.log (response)
            alert ("Saved: " + this.posts.Title)
              })
               window.location = '/'
          },
    getData (){
     // this.id
     var that=this
      axios.get ('http://localhost:8000/posts/'+this.id)
        .then(function (response) {
        that.posts = response.data
        })

    },
    mess: function (message, event) {
     if (event) {
          event.preventDefault()
        }
          alert(message);
          window.location = '/';
         },


  },
  mounted: function (){
    this.getData()

  }

}
</script>

