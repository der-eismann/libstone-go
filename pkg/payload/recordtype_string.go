// Code generated by "stringer -type RecordType,RecordTag pkg/payload/meta.go"; DO NOT EDIT.

package payload

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[RecordTypeUnknown-0]
	_ = x[RecordTypeInt8-1]
	_ = x[RecordTypeUint8-2]
	_ = x[RecordTypeInt16-3]
	_ = x[RecordTypeUint16-4]
	_ = x[RecordTypeInt32-5]
	_ = x[RecordTypeUint32-6]
	_ = x[RecordTypeInt64-7]
	_ = x[RecordTypeUint64-8]
	_ = x[RecordTypeString-9]
	_ = x[RecordTypeDependency-10]
	_ = x[RecordTypeProvider-11]
}

const _RecordType_name = "RecordTypeUnknownRecordTypeInt8RecordTypeUint8RecordTypeInt16RecordTypeUint16RecordTypeInt32RecordTypeUint32RecordTypeInt64RecordTypeUint64RecordTypeStringRecordTypeDependencyRecordTypeProvider"

var _RecordType_index = [...]uint8{0, 17, 31, 46, 61, 77, 92, 108, 123, 139, 155, 175, 193}

func (i RecordType) String() string {
	if i >= RecordType(len(_RecordType_index)-1) {
		return "RecordType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _RecordType_name[_RecordType_index[i]:_RecordType_index[i+1]]
}
func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[RecordTagName-1]
	_ = x[RecordTagArchitecture-2]
	_ = x[RecordTagVersion-3]
	_ = x[RecordTagSummary-4]
	_ = x[RecordTagDescription-5]
	_ = x[RecordTagHomepage-6]
	_ = x[RecordTagSourceID-7]
	_ = x[RecordTagDepends-8]
	_ = x[RecordTagProvides-9]
	_ = x[RecordTagConflicts-10]
	_ = x[RecordTagRelease-11]
	_ = x[RecordTagLicense-12]
	_ = x[RecordTagBuildRelease-13]
	_ = x[RecordTagPackageURI-14]
	_ = x[RecordTagPackageHash-15]
	_ = x[RecordTagPackageSize-16]
	_ = x[RecordTagBuildDepends-17]
	_ = x[RecordTagSourceURI-18]
	_ = x[RecordTagSourcePath-19]
	_ = x[RecordTagSourceRef-20]
}

const _RecordTag_name = "RecordTagNameRecordTagArchitectureRecordTagVersionRecordTagSummaryRecordTagDescriptionRecordTagHomepageRecordTagSourceIDRecordTagDependsRecordTagProvidesRecordTagConflictsRecordTagReleaseRecordTagLicenseRecordTagBuildReleaseRecordTagPackageURIRecordTagPackageHashRecordTagPackageSizeRecordTagBuildDependsRecordTagSourceURIRecordTagSourcePathRecordTagSourceRef"

var _RecordTag_index = [...]uint16{0, 13, 34, 50, 66, 86, 103, 120, 136, 153, 171, 187, 203, 224, 243, 263, 283, 304, 322, 341, 359}

func (i RecordTag) String() string {
	i -= 1
	if i >= RecordTag(len(_RecordTag_index)-1) {
		return "RecordTag(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _RecordTag_name[_RecordTag_index[i]:_RecordTag_index[i+1]]
}