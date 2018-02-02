package DataStruct

// Dictionary ...
type Dictionary struct {
	items map[interface{}]interface{}
}

// Set set dictionary key value
func (d *Dictionary) Set(key, value interface{}) {
	if d.items == nil {
		d.items = make(map[interface{}]interface{})
	}
	d.items[key] = value
}

// Remove remove key
func (d *Dictionary) Remove(key interface{}) bool {
	if d.Has(key) {
		delete(d.items, key)
		return true
	}
	return false
}

// Has Dictionary has key
func (d *Dictionary) Has(key interface{}) bool {
	_, exist := d.items[key]
	return exist
}

// Get ...
func (d *Dictionary) Get(key interface{}) interface{} {
	if d.Has(key) {
		return d.items[key]
	}
	return nil
}

// Clear ...
func (d *Dictionary) Clear() {
	d.items = nil
}

// Size ...
func (d *Dictionary) Size() int {
	cnt := 0
	for range d.items {
		cnt++
	}
	return cnt
}

// Keys ...
func (d *Dictionary) Keys() []interface{} {
	var keys []interface{}
	for k := range d.items {
		keys = append(keys, k)
	}
	return keys
}

// Values ...
func (d *Dictionary) Values() []interface{} {
	var values []interface{}
	for _, v := range d.items {
		values = append(values, v)
	}
	return values
}
