package cluster

import (
	"github.com/zalando/postgres-operator/pkg/spec"
	"testing"

	acidv1 "github.com/zalando/postgres-operator/pkg/apis/acid.zalan.do/v1"
	"github.com/zalando/postgres-operator/pkg/util/config"
	"github.com/zalando/postgres-operator/pkg/util/k8sutil"

	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func TestSyncLogicalBackupJob(t *testing.T){
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
		}, logger, eventRecorder)

	cluster.Statefulset = &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-sts",
		},
	}
	cluster.systemUsers = map[string]spec.PgUser{
		"superuser": spec.PgUser{Origin: spec.RoleOriginInfrastructure},
	}

	clusterMissingObjects := *cluster
	clusterMissingObjects.KubeClient = k8sutil.NewMockKubernetesClient()

	clusterMock := *cluster
	err := clusterMock.syncLogicalBackupJob()
	if  err != nil {
		t.Errorf("Sync PreparedDBs test: Could not synchronize, %+v", err)
	}
}
func TestSyncSecrets(t *testing.T){
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
		}, logger, eventRecorder)

	cluster.Statefulset = &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-sts",
		},
	}
	cluster.systemUsers = map[string]spec.PgUser{
		"superuser": spec.PgUser{Origin: spec.RoleOriginInfrastructure},
	}

	clusterMissingObjects := *cluster
	clusterMissingObjects.KubeClient = k8sutil.NewMockKubernetesClient()

	clusterMock := *cluster
	err := clusterMock.syncSecrets()
	if  err != nil {
		t.Errorf("Sync PreparedDBs test: Could not synchronize, %+v", err)
	}
}

func TestSyncServices(t *testing.T){
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
		}, logger, eventRecorder)

	cluster.Statefulset = &appsv1.StatefulSet{
		ObjectMeta: metav1.ObjectMeta{
			Name: "test-sts",
		},
	}
	cluster.systemUsers = map[string]spec.PgUser{
		"superuser": spec.PgUser{Origin: spec.RoleOriginInfrastructure},
	}

	clusterMissingObjects := *cluster
	clusterMissingObjects.KubeClient = k8sutil.NewMockKubernetesClient()

	clusterMock := *cluster
	err := clusterMock.syncServices()
	if  err != nil {
		t.Errorf("Sync PreparedDBs test: Could not synchronize, %+v", err)
	}
}

func testSyncDatabases(t *testing.T) {
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
				Databases: map[string]string{
					"foo_db": "zalando",
				},
				Users: map[string]acidv1.UserFlags{
					"zalando": acidv1.UserFlags{"superuser", "createdb"},
				},
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
	clusterMissingObjects.KubeClient = k8sutil.NewMockKubernetesClient()

	clusterMock := *cluster
	err := clusterMock.syncDatabases()
	if  err != nil {
		t.Errorf("Sync PreparedDBs test: Could not synchronize, %+v", err)
	}
}

func testSyncPreparedDatabases(t *testing.T) {
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