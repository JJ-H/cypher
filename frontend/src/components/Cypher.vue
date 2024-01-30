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
const pwdVisible = reactive({});

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

const togglePasswordVisibility = (domain) => {
  pwdVisible[domain] = !pwdVisible[domain] ?? false;
  if (pwdVisible[domain]) {
    setTimeout(() => {
      pwdVisible[domain] = false;
    }, 1000 * 10);
  }
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
            <div class="el-custom-password-column" v-else>
              <span class="ellipsis">{{
                pwdVisible[scope.row.domain] ? scope.row.password : "******"
              }}</span>
              <span
                class="el-custom-hover-icon"
                @click="togglePasswordVisibility(scope.row.domain)"
              >
                <el-icon v-if="!pwdVisible[scope.row.domain]">
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 1024 1024"
                    data-v-ea893728=""
                  >
                    <path
                      fill="currentColor"
                      d="M512 160c320 0 512 352 512 352S832 864 512 864 0 512 0 512s192-352 512-352m0 64c-225.28 0-384.128 208.064-436.8 288 52.608 79.872 211.456 288 436.8 288 225.28 0 384.128-208.064 436.8-288-52.608-79.872-211.456-288-436.8-288zm0 64a224 224 0 1 1 0 448 224 224 0 0 1 0-448m0 64a160.192 160.192 0 0 0-160 160c0 88.192 71.744 160 160 160s160-71.808 160-160-71.744-160-160-160"
                    ></path>
                  </svg>
                </el-icon>
                <el-icon v-else>
                  <svg
                    xmlns="http://www.w3.org/2000/svg"
                    viewBox="0 0 1024 1024"
                    data-v-ea893728=""
                  >
                    <path
                      fill="currentColor"
                      d="M876.8 156.8c0-9.6-3.2-16-9.6-22.4-6.4-6.4-12.8-9.6-22.4-9.6-9.6 0-16 3.2-22.4 9.6L736 220.8c-64-32-137.6-51.2-224-60.8-160 16-288 73.6-377.6 176C44.8 438.4 0 496 0 512s48 73.6 134.4 176c22.4 25.6 44.8 48 73.6 67.2l-86.4 89.6c-6.4 6.4-9.6 12.8-9.6 22.4 0 9.6 3.2 16 9.6 22.4 6.4 6.4 12.8 9.6 22.4 9.6 9.6 0 16-3.2 22.4-9.6l704-710.4c3.2-6.4 6.4-12.8 6.4-22.4Zm-646.4 528c-76.8-70.4-128-128-153.6-172.8 28.8-48 80-105.6 153.6-172.8C304 272 400 230.4 512 224c64 3.2 124.8 19.2 176 44.8l-54.4 54.4C598.4 300.8 560 288 512 288c-64 0-115.2 22.4-160 64s-64 96-64 160c0 48 12.8 89.6 35.2 124.8L256 707.2c-9.6-6.4-19.2-16-25.6-22.4Zm140.8-96c-12.8-22.4-19.2-48-19.2-76.8 0-44.8 16-83.2 48-112 32-28.8 67.2-48 112-48 28.8 0 54.4 6.4 73.6 19.2zM889.599 336c-12.8-16-28.8-28.8-41.6-41.6l-48 48c73.6 67.2 124.8 124.8 150.4 169.6-28.8 48-80 105.6-153.6 172.8-73.6 67.2-172.8 108.8-284.8 115.2-51.2-3.2-99.2-12.8-140.8-28.8l-48 48c57.6 22.4 118.4 38.4 188.8 44.8 160-16 288-73.6 377.6-176C979.199 585.6 1024 528 1024 512s-48.001-73.6-134.401-176Z"
                    ></path>
                    <path
                      fill="currentColor"
                      d="M511.998 672c-12.8 0-25.6-3.2-38.4-6.4l-51.2 51.2c28.8 12.8 57.6 19.2 89.6 19.2 64 0 115.2-22.4 160-64 41.6-41.6 64-96 64-160 0-32-6.4-64-19.2-89.6l-51.2 51.2c3.2 12.8 6.4 25.6 6.4 38.4 0 44.8-16 83.2-48 112-32 28.8-67.2 48-112 48Z"
                    ></path>
                  </svg>
                </el-icon>
              </span>
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
        <el-table-column
          label="操作"
          class="el-custom-operation-column"
          width="135px"
          align="center"
          :show-overflow-tooltip="false"
        >
          <template #default="scope">
            <span class="el-custom-flex" v-if="scope.row.__add__">
              <el-space>
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
            </span>
            <template v-else>
              <span class="el-custom-flex">
                <el-space>
                  <el-button
                    size="small"
                    @click="copyPassword(scope.row.password)"
                    >复制</el-button
                  >
                  <el-popconfirm
                    hide-icon
                    title="确认删除吗？"
                    confirm-button-text="确认"
                    cancel-button-text="取消"
                    @confirm="handleDelete(scope.row.domain)"
                  >
                    <template #reference>
                      <el-button size="small" type="danger">删除</el-button>
                    </template>
                  </el-popconfirm>
                </el-space>
              </span>
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

main {
  position: fixed;
  height: 100%;
  width: 100%;
  overflow-y: scroll;
}

.fixed {
  position: sticky;
  z-index: 100;
  top: 0;
}
.el-custom-form {
  margin-top: 8px;
}
.el-custom-password-column {
  overflow: hidden;
  text-overflow: ellipsis;
  padding-right: 34px;
}
.el-custom-password-column .ellipsis {
  text-overflow: ellipsis;
  overflow: hidden;
  white-space: nowrap;
}

.el-custom-password-column:hover .el-custom-hover-icon {
  display: block;
}
.el-custom-hover-icon {
  display: none;
  position: absolute;
  top: 50%;
  transform: translateY(-50%);
  right: 20px;
  cursor: pointer;
}
.el-custom-form :deep(.el-form-item) {
  margin-bottom: 0;
}
.el-custom-flex {
  display: flex;
}
</style>
