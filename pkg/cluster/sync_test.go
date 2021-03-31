package cluster

import (
	"testing"

	acidv1 "github.com/zalando/postgres-operator/pkg/apis/acid.zalan.do/v1"
	"github.com/zalando/postgres-operator/pkg/util/config"
	"github.com/zalando/postgres-operator/pkg/util/k8sutil"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)


func TestSyncPreparedDatabases(t *testing.T) {
	var cluster = New(
		Config{
			OpConfig: config.Config{
				ProtectedRoles: []string{"admin"},
				Auth: config.Auth{
					SuperUsername:       superUserName,
					ReplicationUsername: replicationUserName,
				},

			},
		}, k8sutil.KubernetesClient{}, acidv1.Postgresql{
			Spec: acidv1.PostgresSpec{
				PreparedDatabases: map[string]acidv1.PreparedDatabase{
					"foo": {
						DefaultUsers: true,
						PreparedSchemas: map[string]acidv1.PreparedSchema{
							"bar": {
								DefaultUsers: true,
							},
						},
						Extensions: map[string]string{
							"pg_partman": "public",
							"pgcrypto": "public",
						},
					},
				},
			},
		}, logger, eventRecorder)

	cluster.Statefulset = &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-sts",
		},
	}

	clusterMissingObjects := *cluster
	clusterMissingObjects.KubeClient = k8sutil.ClientMissingObjects()

	clusterMock := *cluster
	err := clusterMock.syncPreparedDatabases()
	if  err != nil {
		t.Errorf("Sync PreparedDBs test: Could not synchronize, %+v", err)
	}
}