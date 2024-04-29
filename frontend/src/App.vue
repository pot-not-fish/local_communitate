<script lang="ts" setup>
  import RecieveFile from './components/RecieveFile.vue'
  import UploadFile from './components/UploadFile.vue'
  import HelloWorld from './components/HelloWorld.vue'
</script>

<template>
  <div class="common-layout">
    <el-container>
      <el-header>
        <h1  style="text-align: center; color: black;">局域网文件传输工具</h1>
      </el-header>

      <el-main>
        <div class="separator"></div>
        <div class="container">
          <div class="box">
            <UploadFile/>
          </div>
          <div class="box">
            <RecieveFile/>
          </div>
        </div>
      </el-main>
    </el-container>
  </div>
</template>

<style>
.separator {  
  height: 30px; /* 设置间隔器的高度 */  
  width: 100%;  
  background-color: transparent; /* 间隔器背景色设置为透明 */  
}  

.container {  
  display: flex;  
  justify-content: space-between;  
} 

.box {  
  width: 47%;  
  height: 100%; /* 可以根据需要调整高度 */   
}

body {  
  background-color: white;  
}
</style>
