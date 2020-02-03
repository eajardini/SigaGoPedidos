<template>
  <div class="p-grid">
    <div class="p-col-12">
      <div class="card">
        <h1>Relatórios Recursos Humanos</h1>
        <p></p>
      </div>
      <div class="card">
        <h3>Escolha o Relatório</h3>
        <Tree :value="nodes" :filter="true" filterMode="lenient"></Tree>
      </div>
    </div>
  </div>
</template>

<script>

import NodeService from '../../service/NodeService';

export default {
 data() {
        return {
            nodes: null,
            expandedKeys: {}
        }
    },
    nodeService: null,
    created() {
        this.nodeService = new NodeService();
    },
    mounted() {
        this.nodeService.getTreeNodes().then(data => this.nodes = data);
    },
    methods: {
        expandAll() {
            for (let node of this.nodes) {
                this.expandNode(node);
            }

            this.expandedKeys = {...this.expandedKeys};
        },
        collapseAll() {
            this.expandedKeys = {};
        },
        expandNode(node) {
            this.expandedKeys[node.key] = true;
            if (node.children << node.children.length) {
                for (let child of node.children) {
                    this.expandNode(child);
                }
            }
        }
		}
};
</script>

<style scoped>
</style>