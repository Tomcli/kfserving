import pytest
import kfserving

def test_storage_local_path():
    abs_path = 'file:///tmp/file'
    relative_path = 'file://.'
    assert kfserving.Storage.download(abs_path) == abs_path.replace("file://", "", 1)
    assert kfserving.Storage.download(relative_path) == relative_path.replace("file://", "", 1)

    # Warning test
    not_exist_path = 'file:///some/random/path'
    assert kfserving.Storage.download(not_exist_path) == not_exist_path.replace("file://", "", 1)
