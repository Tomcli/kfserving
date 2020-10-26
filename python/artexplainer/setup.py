# Copyright 2019 kubeflow.org.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#    http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

from setuptools import setup, find_packages

tests_require = [
    'pytest',
    'pytest-tornasync',
    'mypy'
]
setup(
    name='artserver',
    version='0.2.1',
    author_email='Andrew.Butler@ibm.com',
    license='https://github.com/kubeflow/kfserving/LICENSE',
    url='https://github.com/kubeflow/kfserving/python/artserver',
    description='Model Server implementation for AI Robustness Toolbox. \
                 Not intended for use outside KFServing Frameworks Images',
    long_description='Model Server implementation for AI Robustness Toolbox. \
                 Not intended for use outside KFServing Frameworks Images',
    python_requires='>3.7',
    packages=find_packages("artserver"),
    install_requires=[
        "kfserving>=0.4.0",
        "argparse >= 1.4.0",
        "numpy >= 1.8.2",
        "tensorflow == 1.14.0",
        "keras >= 2.3.1",
        "adversarial-robustness-toolbox == 1.4.1",
        "nest_asyncio>=1.4.0",
        "kornia==0.4.0"
    ],
    tests_require=tests_require,
    extras_require={'test': tests_require}
)
