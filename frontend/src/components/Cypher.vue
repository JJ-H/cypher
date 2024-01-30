<script setup>
import { ElMessage } from "element-plus";
import { reactive, ref, onMounted, computed } from "vue";
import {
  ListCredential,
  SetCredential,
  DeleteCypherByDomain,
} from "../../wailsjs/go/services/CredentialService";

const defaultFormValue = {
  domain: "",
  username: "",
  password: "",
  note: "",
};
const data = reactive({
  loading: ref(false),
  cyphers: ref([]),
  keyword: ref(""),
  selectType: ref(""),
  showPassword: ref(false),
  createOrEditDialogVisible: ref(false),
});

const formRef = ref();
const formRules = reactive({
  domain: [
    {
      required: true,
      message: "请输入域名",
      trigger: "change",
    },
  ],
  username: [
    {
      required: true,
      message: "请输入用户名",
      trigger: "change",
    },
  ],
  password: [
    {
      required: true,
      message: "请输入密码",
      trigger: "change",
    },
  ],
});

const form = reactive(defaultFormValue);

const togglePasswordVisibility = () => {
  data.showPassword = !data.showPassword;
};

const filterCyphers = computed(() =>
  data.cyphers.filter((cypher) => {
    if (cypher.__add__) {
      return true;
    }
    var key = "note";
    if (data.selectType == "1") {
      key = "domain";
    } else if (data.selectType == "2") {
      key = "username";
    }
    const keywordMatch =
      !data.keyword ||
      cypher[key].toLowerCase().includes(data.keyword.toLowerCase());

    return keywordMatch;
  })
);
const listCredential = () => {
  data.loading = true;
  return new Promise(() => {
    ListCredential()
      .then((result) => {
        data.cyphers = result;
      })
      .finally(() => {
        data.loading = false;
      });
  });
};

const copyPassword = (password) => {
  navigator.clipboard
    .writeText(password)
    .then(() => {
      console.log("Password copied successfully:", password);
    })
    .catch((error) => {
      console.error("Error copying password:", error);
    });
};

const handleAdd = () => {
  if (!data.cyphers.some((o) => o.__add__)) {
    data.cyphers = [
      {
        __add__: true,
        domain: "",
        username: "",
        password: "",
        note: "",
      },
      ...data.cyphers,
    ];
  }
};

const handleAddCancel = (formEl) => {
  if (!formEl) return;
  formEl.resetFields();
  data.cyphers = data.cyphers.filter((o) => !o.__add__);
};

const handleAddConfirm = async (formEl) => {
  if (!formEl) return;

  await formEl.validate(async (valid, fields) => {
    if (valid) {
      if (data.cyphers.some((o) => o.domain === form.domain)) {
        ElMessage.error("已存在相同域名");
        return;
      }
      await SetCredential(form);
      handleAddCancel(formEl);
      await listCredential();
    } else {
      console.log("error submit!", fields);
    }
  });
};

const handleDelete = async (domain) => {
  await DeleteCypherByDomain(domain);
  await listCredential();
};

onMounted(() => {
  // 主动加载 cyphers
  listCredential();
});
</script>

<template>
  <main>
    <div class="flex-container fixed">
      <el-text class="mx-1 logo" type="success">🪺Cypher</el-text>
      <div class="mt-4 flex-container">
        <el-space :size="8">
          <el-input
            v-model="data.keyword"
            placeholder="请输入关键字进行搜索"
            class="input-with-select"
          >
            <template #append>
              <el-select
                v-model="data.selectType"
                default-first-option
                placeholder="备注"
                style="width: 115px"
              >
                <el-option label="域名" value="1" />
                <el-option label="用户" value="2" />
                <el-option label="备注" value="3" />
              </el-select>
            </template>
          </el-input>
          <el-button
            class="visible-btn"
            size="middle"
            @click="togglePasswordVisibility"
            >密明文切换</el-button
          >
          <el-button type="primary" @click="handleAdd"> 添加 </el-button>
        </el-space>
      </div>
    </div>
    <el-form
      hide-required-asterisk
      label-position="left"
      label-width="0"
      :model="form"
      :rules="formRules"
      class="el-custom-form"
      ref="formRef"
    >
      <el-table
        v-if="data.cyphers"
        v-loading="data.loading"
        :data="filterCyphers"
        border
        stripe
        highlight-current-row
        style="width: 100%"
        v-load
        empty-text="暂无数据"
      >
        <el-table-column prop="domain" label="域名" width="180" align="left">
          <template #default="scope">
            <el-form-item v-if="scope.row.__add__" prop="domain">
              <el-input v-model="form.domain" placeholder="请输入域名" />
            </el-form-item>
            <span v-else>{{ scope.row.domain }}</span>
          </template>
        </el-table-column>
        <el-table-column
          prop="username"
          label="用户名"
          width="180"
          align="left"
        >
          <template #default="scope">
            <el-form-item v-if="scope.row.__add__" prop="username">
              <el-input v-model="form.username" placeholder="请输入用户名" />
            </el-form-item>
            <span v-else>{{ scope.row.username }}</span>
          </template>
        </el-table-column>
        <el-table-column label="密码" align="left">
          <template #default="scope">
            <el-form-item v-if="scope.row.__add__" prop="password">
              <el-input
                v-model="form.password"
                type="password"
                placeholder="请输入密码"
                show-password
              />
            </el-form-item>
            <div v-else>
              <span v-if="data.showPassword" disabled>{{
                scope.row.password
              }}</span>
              <span v-else>******</span>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="note" label="备注" align="left">
          <template #default="scope">
            <el-form-item v-if="scope.row.__add__" prop="note">
              <el-input
                rows="1"
                v-model="form.note"
                type="textarea"
                placeholder="请输入备注"
              />
            </el-form-item>
            <span v-else>{{ scope.row.note }}</span>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="130px" align="center">
          <template #default="scope">
            <el-space v-if="scope.row.__add__">
              <el-button
                size="small"
                type="primary"
                @click="handleAddConfirm(formRef)"
                >保存</el-button
              >
              <el-button size="small" @click="handleAddCancel(formRef)"
                >取消</el-button
              >
            </el-space>
            <template v-else>
              <el-space>
                <el-button
                  size="small"
                  @click="copyPassword(scope.row.password)"
                  >复制</el-button
                >
                <el-button
                  size="small"
                  type="danger"
                  @click="handleDelete(scope.row.domain)"
                  >删除</el-button
                >
              </el-space>
            </template>
          </template>
        </el-table-column>
      </el-table>
    </el-form>
  </main>
</template>

<style scoped>
.logo {
  font-size: 30px;
}
.flex-container {
  padding-left: 10px;
  padding-right: 5px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background: white;
}

.visible-btn {
  color: #81cbf5;
  background-color: #f8f8f8;
}
main {
  position: fixed;
  height: 100%;
  width: 100%;
  overflow-y: scroll;
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
}

.fixed {
  position: sticky;
  z-index: 100;
  top: 0;
}
.el-custom-form {
  -webkit-user-select: none;
  -moz-user-select: none;
  -ms-user-select: none;
  user-select: none;
  margin-top: 8px;
}
.el-custom-form ::v-deep .el-form-item {
  margin-bottom: 0;
}
</style>
