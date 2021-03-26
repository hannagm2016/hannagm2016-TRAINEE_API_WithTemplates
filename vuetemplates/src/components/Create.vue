<template>
<div>
    <div class="row py-lg-5">
                <div class="col-lg-6 col-md-8 mx-auto">
                    <form action="/savePost" method="post" >
                        <input type="text" name="title" placeholder="Title of your post" class="form-control"><br>
                        <textarea name="body" id="body"  placeholder="Body of your post" class="form-control"></textarea><br>
                        <button class="btn btn-primary my-2">Create</button>
                    </form>
                </div>

            </div>

  </div>
</template>

<script>
export default {

}
</script>

