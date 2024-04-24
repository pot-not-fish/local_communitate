<script lang="ts" setup>
  import ReceiveFile from './components/ReceiveFile.vue'
  import UploadFile from './components/UploadFile.vue'
</script>

<template>
  <div class="common-layout">
    <el-container>
      <el-header>
        <h1 style="text-align: center">局域网文件传输工具</h1>
      </el-header>

      <el-main>
        <div class="container">
          <div class="box">
            <UploadFile/>
          </div>
          <div class="box">
            <ReceiveFile/>
          </div>
        </div>
      </el-main>
    </el-container>
  </div>
</template>

<style>
.container {  
  display: flex;  
  justify-content: space-between;  
} 

.box {  
  width: 50%;  
  height: 100%; /* 可以根据需要调整高度 */   
}
</style>
