<template>
<div>
  <h1>Edit {{id}}</h1>

    <div class="row py-lg-5">
                <div class="col-lg-6 col-md-8 mx-auto">
                    <form action="http://localhost:8000/savePost" method="post" >
                        <input type="hidden" name="id" :value="post.Id" />
                        <input type="hidden" name="userid" :value="post.UserId" />
                        <input type="text" name="title" placeholder="Title of your post" class="form-control" :value="post.Title"><br>
                        <textarea name="body" id="body"  placeholder="Body of your post" class="form-control" :value="post.Body"></textarea><br>
                        <button class="btn btn-primary my-2" >Save</button>
                    </form>
                    <p>{{post.Id}}</p>
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
      axios.get ('http://localhost:8000/posts/'+this.id)
        .then(function (response) {
        that.post = response.data
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

