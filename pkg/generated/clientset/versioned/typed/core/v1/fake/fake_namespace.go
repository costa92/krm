
// Finalize takes the representation of a namespace to update.  Returns the server's representation of the namespace, and an error, if it occurs.
func (c *FakeNamespaces) Finalize(ctx context.Context, namespace *v1.Namespace, opts metav1.UpdateOptions) (result *v1.Namespace, err error) {
	return nil, nil
}
