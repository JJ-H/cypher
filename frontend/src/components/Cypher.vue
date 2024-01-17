<script setup>
import {reactive, ref, onMounted, computed} from 'vue'
import {ListCredential} from "../../wailsjs/go/services/CredentialService";


const data = reactive({
  cyphers: ref(null),
  keyword: ref(""),
  selectType: ref("")
})

const filterCyphers = computed(() =>
    data.cyphers.filter((cypher) => {
      var key = "note"
      if (data.selectType == "1") {
        key = "domain"
      } else if(data.selectType == "2") {
        key = "username"
      }
      const keywordMatch =
          !data.keyword ||
          cypher[key].toLowerCase().includes(data.keyword.toLowerCase());

      return keywordMatch;
    })
);

onMounted(() => {
  // 主动加载 cyphers
  listCredential();
});

function listCredential() {
  ListCredential()
      .then(result => {
        data.cyphers = result
      })
}

function copyPassword(password) {
  // Use Clipboard API or a library like clipboard.js to copy the password
  // For simplicity, let's use the Clipboard API
  navigator.clipboard.writeText(password).then(() => {
    // Password copied successfully
    // You can also show a notification or perform any other action
    console.log('Password copied successfully:', password);
  }).catch((error) => {
    // Handle clipboard write error
    console.error('Error copying password:', error);
  });
}

</script>

<template>
  <main>
    <div class="flex-container">
      <el-text class="mx-1 logo" type="success">🪺Cypher</el-text>

      <div class="mt-4">
        <el-input
            v-model="data.keyword"
            placeholder="请输入关键字进行搜索"
            class="input-with-select"
        >
          <template #append>
            <el-select v-model="data.selectType" default-first-option placeholder="备注" style="width: 115px">
              <el-option label="域名" value="1" />
              <el-option label="用户" value="2" />
              <el-option label="备注" value="3" />
            </el-select>
          </template>
        </el-input>
      </div>
    </div>
    <el-table v-if="data.cyphers" :data="filterCyphers" border stripe highlight-current-row style="width: 100%">
      <el-table-column prop="domain" label="域名" width="180" align="left"/>
      <el-table-column prop="username" label="用户名" width="180" align="left" />
      <el-table-column prop="password" label="密码" align="left" />
      <el-table-column prop="note" label="备注" align="left" />
      <el-table-column label="操作" width="80px" align="center">
        <template #default="scope">
          <el-button size="small" @click="copyPassword(scope.row.password)">复制</el-button>
<!--          <el-button size="small" type="danger" >删除</el-button>-->
        </template>
      </el-table-column>
    </el-table>
  </main>
</template>

<style scoped>
  .logo {
    font-size: 30px;
  }
  .flex-container {
    padding-left: 10px;
    padding-right: 10px;
    display: flex;
    justify-content: space-between;
    align-items: center;
    background: white;
  }
</style>
