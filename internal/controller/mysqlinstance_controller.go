package controller

import (
	"context"
	"fmt"

	"github.com/pingcap/errors"
	appsv1 "k8s.io/api/apps/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/log"

	mysqlv1 "github.com/CRDProject/api/v1"
)

// MySQLInstanceReconciler reconciles a MySQLInstance object
type MySQLInstanceReconciler struct {
	client.Client
	Scheme *runtime.Scheme
}

//+kubebuilder:rbac:groups=mysql.my.domain,resources=mysqlinstances,verbs=get;list;watch;create;update;patch;delete
//+kubebuilder:rbac:groups=mysql.my.domain,resources=mysqlinstances/status,verbs=get;update;patch
//+kubebuilder:rbac:groups=mysql.my.domain,resources=mysqlinstances/finalizers,verbs=update

// Reconcile is part of the main kubernetes reconciliation loop which aims to
// move the current state of the cluster closer to the desired state.
// TODO(user): Modify the Reconcile function to compare the state specified by
// the MySQLInstance object against the actual cluster state, and then
// perform operations to make the cluster state reflect the state specified by
// the user.
//
// For more details, check Reconcile and its Result here:
// - https://pkg.go.dev/sigs.k8s.io/controller-runtime@v0.15.0/pkg/reconcile
func (r *MySQLInstanceReconciler) Reconcile(ctx context.Context, req ctrl.Request) (ctrl.Result, error) {
	_ = log.FromContext(ctx)
	// Step 1: Fetch the MySQLInstance instance
	mysqlInstance := &mysqlv1.MySQLInstance{}
	err := r.Get(ctx, req.NamespacedName, mysqlInstance)
	if err != nil {
		if errors.IsNotFound(err) {
			// Object not found, return. Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return ctrl.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return ctrl.Result{}, err
	}

	// Step 2: Check if the MySQL instance already exists
	// For the sake of this example, let's assume that we are using a dummy function 'CheckIfMySQLInstanceExists'
	// which checks if a MySQL instance exists based on the provided MySQLInstance spec.
	exists, err := r.CheckIfMySQLInstanceExists(ctx, req.Namespace, mysqlInstance.Name)
	if err != nil {
		// If an error occurs during checking, requeue the request.
		return ctrl.Result{}, err
	}

	// If the MySQL instance does not exist, create it.
	if !exists {
		err := r.CreateMySQLInstance(ctx, mysqlInstance)
		if err != nil {
			// If an error occurs during MySQL instance creation, requeue the request.
			return ctrl.Result{}, err
		}
		// MySQL instance successfully created.
		mysqlInstance.Status.Conditions = append(mysqlInstance.Status.Conditions, metav1.Condition{
			Type:   "Created",
			Status: metav1.ConditionTrue,
		})
		r.Status().Update(ctx, mysqlInstance)
		fmt.Println("MySQL instance created successfully")
	}

	// TODO(user): your logic here

	return ctrl.Result{}, nil
}

// SetupWithManager sets up the controller with the Manager.
func (r *MySQLInstanceReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&mysqlv1.MySQLInstance{}).
		Complete(r)
}

func (r *MySQLInstanceReconciler) CheckIfMySQLInstanceExists(ctx context.Context, namespace string, name string) (bool, error) {
	statefulSet := &appsv1.StatefulSet{}
	err := r.Client.Get(ctx, client.ObjectKey{Namespace: namespace, Name: name}, statefulSet)
	if err != nil {
		if errors.IsNotFound(err) {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func (r *MySQLInstanceReconciler) CreateMySQLInstance(ctx context.Context, mysqlInstance *mysqlv1.MySQLInstance) error {
	statefulSet := &appsv1.StatefulSet{ /* 填充 StatefulSet 对象的详细信息 */ }
	err := r.Client.Create(ctx, statefulSet)
	return err
}
